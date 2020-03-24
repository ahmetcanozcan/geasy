const express = require('express');

const app = express();


const port = process.argv[2]

console.log('DB module started to port ', port);
const users = [
  {
    username: "excalibur12",
    password: "passwordispassword",
    firstname: "adam",
    lastname: "sandler"
  }
]


app.get('/:username/:password', function (req, res) {
  var username = req.params["username"];
  var password = req.params["password"];
  var user = users.filter(user => user.username == username && user.password == password);
  if (user.length == 0) {
    return res.send({ error: "User Not Found" });
  }
  return res.send(user[0]);
});



app.listen(port, () => console.log('DB module listening to port: ', port));