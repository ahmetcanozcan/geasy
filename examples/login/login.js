const username = process.argv[2];
const password = process.argv[3];
const dbPort = process.argv[4].trim();

const axios = require('axios').default;

axios.get(`http://localhost:${dbPort}/${username}/${password}`)
  .then(response => response.data)
  .then(response => {
    if (response.error) {
      return console.log(response);
    }
    Object.assign(response, { status: "Ok!" });
    console.log(response);
  }).catch(console.log);