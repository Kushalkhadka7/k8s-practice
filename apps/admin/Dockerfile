# Stage: Development
FROM node:erbium-slim as dev

# Make working directory app/
WORKDIR /app

# Copy pacakges files.
COPY package.json /app/
COPY yarn.lock /app/

# Install app packages.
RUN yarn

# Copy source code.
COPY . /app/

# Expose post to run the application.
EXPOSE 3000

CMD [ "yarn", "start" ]

# Stage: Builder
FROM node:erbium-slim as builder

WORKDIR /app

COPY --from=dev /app /app

# [NOTE]: This needs to be replace by yarn build.
RUN yarn

# Stage: Production
FROM node:erbium-slim as prod

WORKDIR /app

COPY --from=builder /app/dist /app/dist

EXPOSE 3000

CMD ["node", "dist/index.js"]