const SERVER_PORT = 3000;
const express = require('express');
const mysql = require('mysql2');
const faker = require('faker');
const app = express();

const persitRandomPeople = (connection ,cb) => {
  const randomName = faker.name.findName();
  const sql = 'INSERT INTO people (name) VALUES (?)';
  connection.query(sql, [randomName], cb);
};

const getPersistedPeoples = (connection, cb) => {
  connection.query('SELECT * FROM people', cb);
};

const getResponse = (res, connection, data) => {
  res.write(data);
  res.end();
  connection.end();
};

app.get('/', (req, res) => {
  const connection = mysql.createConnection({
    host: 'db',
    user: 'root',
    password: 'secret',
    database: 'homestead'
  });

  persitRandomPeople(connection, (err) => {
    if (err) {
      return getResponse(res, connection, JSON.stringify(err));
    }

    getPersistedPeoples(connection, (err, peoples) => {
      if (err) {
        return getResponse(res, connection, JSON.stringify(err));
      }

      const data = `
        <h1>Full Cycle Rocks!</h1>
        <hr>
        <ul>
          ${peoples.map((people) => (
            `<li>${people.name}</li>`
          )).join('')}
        </ul>
      `;

      return getResponse(res, connection, data);
    });
  });
})

app.listen(SERVER_PORT, () => {
  console.log(`Example app listening at ${SERVER_PORT}`)
});
