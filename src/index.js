var axios = require("axios")
var button= document.querySelector('button');
var input=document.querySelector('#password')
var status = document.querySelector("#status");
var errors = document.querySelector('#errors')
console.log(button);

button.addEventListener('click',function (){
    while(errors.hasChildNodes()){
        errors.removeChild(errors.lastChild)
    }
    var inp =input.value;
    axios.post("http://localhost:8080",{
        value: inp
    })
        .then(function (response){
        var stat = response.data;
        status.innerText = stat.message;
        stat.errors.forEach(function(element) {
            var error =document.createElement('p')
            error.innerText=element;
            errors.appendChild(error)
        });

        })
        .catch(function (err){
            console.log(err)
        })


    
    })