import got from 'got';
import cors from 'cors';
import helmet from 'helmet';
import dotenv from 'dotenv';
import express from 'express';
import bodyParser from 'body-parser';
import compression from 'compression';

const app = express();

dotenv.config();

app.set('PORT', process.env.PORT);
app.set('NAME', process.env.NAME);
app.set('HOST', process.env.HOST);

app.use(cors());
app.use(helmet());
app.use(compression());
app.use(bodyParser.json());

app.get('/info', (req, res, next) => {
  res.send({
    name: 'admin-api',
    version: '1.0.0',
  });
});

app.get('/admin', async (req, res, next) => {
  const url = process.env.AUTH_URL + 'login';
  const response = await got.get(url);

  res.send(response.body);
});

app.listen(app.get('PORT') || 3000, app.get('host') || 'localhost', () => {
  console.log(`Server started at http://localhost:${app.get('PORT')}`);
});
