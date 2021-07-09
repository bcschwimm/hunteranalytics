/*
work in progress, todo: place .behavior-data in html
*/
behaviorTable = document.querySelector("..behavior-data")
            fetch("/behaviorapi")
            .then(response => response.json())
            .then(behaviorRows => {
                behaviorRows.forEach(elm => {
                    row = document.createElement("tr")

                    date = document.createElement("td")
                    date.innerHTML = elm.date

                    crate = document.createElement("td")
                    crate.innerHTML = elm.crate

                    notes = document.createElement("td")
                    notes.innerHTML = elm.notes

                    row.appendChild(date)
                    row.appendChild(crate)
                    row.appendChild(notes)

                    behaviorTable.appendChild(row)
                })
            })