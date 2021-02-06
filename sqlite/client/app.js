let table = document.querySelector("table");
fetch('http://localhost:8080/api/v1/books', {})
    .then(res => res.json())
    .then(data => {
        let books = data
        let obj = Object.keys(books[0])
        generateTableHead(table, obj)
        generateTable(table, books)
    }).catch(err => {
        console.log(err)
    })

// let main = document.getElementById('datajs')

function generateTableHead(table, data) {
    let thead = table.createTHead()
    let row = thead.insertRow()
    for (let key of data) {
        let th = document.createElement("th")
        let text = document.createTextNode(key)
        th.appendChild(text)
        row.appendChild(th)
    }
}
function generateTable(table, data) {
    for (let element of data) {
        let row = table.insertRow()
        for (key in element) {
            let cell = row.insertCell()
            let text = document.createTextNode(element[key])
            cell.appendChild(text)
        }
    }
}

const myForm = document.getElementById('form')
myForm.addEventListener('submit', async (e) => {
    e.preventDefault()
    const formData = new FormData(myForm)
    // console.log(formData); -> FormData { }
    const jsonObj = Object.fromEntries(formData) // method transforms a list of key-value pairs into an object
    // console.log(jsonObj); -> Object {name: "awa", author: "awsd"}
    try {
        const response = await fetch('http://localhost:8080/api/v1/books', {
            method: 'POST',
            body: JSON.stringify(jsonObj),
            headers: {
                'Content-Type': 'application/json; charset=utf-8'
            }
        })
        const json = await response.json();
        console.log(json);
        myForm.reset()
        window.location.reload()
    } catch (e) {
        console.error(e);
    }
})
