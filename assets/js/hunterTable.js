hunterTable = document.querySelector(".hunter-data")
            fetch("/api")
            .then(response => response.json())
            .then(hunterRows => {
                hunterRows.forEach(elm => {
                    // create table row
                    row = document.createElement("tr")

                    // create the table data elements
                    playing = document.createElement("td")
                    playing.innerHTML = elm.playing

                    training = document.createElement("td")
                    training.innerHTML = elm.training

                    exercising = document.createElement("td")
                    exercising.innerHTML = elm.exercising

                    woofing = document.createElement("td")
                    woofing.innerHTML = elm.woofing

                    date = document.createElement("td")
                    date.innerHTML = elm.date

                    // add data elms to row
                    row.appendChild(date)
                    row.appendChild(playing)
                    row.appendChild(training)
                    row.appendChild(exercising)
                    row.appendChild(woofing)

                    hunterTable.appendChild(row)
                })
            })