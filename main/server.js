const express = require('express');
const path = require("path");
const axios = require("axios");
const app = express();

const port = 4000;

app.use(express.json());
app.use(express.urlencoded({extended: false}));

app.disable('x-powered-by');

// CORS Header
app.use((req, res, next) => {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE");
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");

    res.header(
        'Access-Control-Expose-Headers',
        'x-access-token'
    );

    next();
});

app.get('/archive/:name', (req, res) => {
    axios.get(process.env.SERVICE_URL + '/archive/' + req.params.name)
        .then(data => {
            const d = String(data.data).replaceAll("\n", "<br/>")
            res.send(d)
        })
        .catch(() => {
            res.sendStatus(404)
        })
})

app.post('/archive', (req, res) => {
    axios.post(process.env.SERVICE_URL + '/archive', req.body)
        .then(data => {
            res.send(data.data)
        })
        .catch(() => {
            res.sendStatus(404)
        })
})

app.use('/', express.static(path.resolve('public')))

app.listen(port, () => {
    console.log(`The server is started on port: ${port}`);
})
