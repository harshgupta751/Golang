const express = require('express')
const cors = require('cors')

const app = express()

app.use(cors())

app.use(express.json())
app.use(express.urlencoded({extended : false}))

app.get('/get', (req, res)=>{

res.json({
message: "Get request successfully executed."
})

})

app.post('/post', (req, res)=>{

res.json(req.body);

})


app.post('/postform', (req, res)=>{


res.send(JSON.stringify(req.body))


})

app.listen(5000, ()=>console.log("Server is running on Port 5000"))