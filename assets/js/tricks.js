/*
work in progress, why is the trickname returning something thats not a string
*/
tricksTable = document.querySelector(".trick-data")
            fetch("/tricksapi")
            .then(response => response.json())
            .then(tricksRows => {
                tricksRows.forEach(elm => {
                    row = document.createElement("tr")

                    trickname = document.createElement("td")
                    trickname.innerHTML = elm.trickname

                    detail = document.createElement("td")
                    detail.innerHTML = elm.detail

                    level = document.createElement("td")
                    level.innerHTML = elm.level

                    row.appendChild(trickname)
                    row.appendChild(level)
                    row.appendChild(detail)

                    tricksTable.appendChild(row)
                })
            })