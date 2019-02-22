var express = require("express");
var bodyParser = require("body-parser");
var app = express();

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.use(express.static("public"));

const port = 5466;

var i = {
  ID_345: {
    formula : "2 * 2 * 2 * 2",
    completed: false
  }
}

app.get("/cal.js", function(req, res) {
  res.send(createScript("ID_345", i.ID_345))
});

app.post("/cal", function(req, res) {
  i = { ...i, [req.body.id]: { ...i[req.body.id], completed: true, answer: req.body.answer}}
  console.log(JSON.stringify(i));
});

app.listen(port, function() {
  console.log(`Example app listening on port ${port}!`);
});

function createScript(id, toEval) {
  return `
  (function() {
    axios
      .post("/cal", {
        id: "${id}",
        answer: eval(${toEval.formula})
      })
  })();
  `;
}