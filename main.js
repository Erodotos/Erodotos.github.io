
function calculate() {

    var baseCharge = 30
    var exchangeRate = 1.7086

    catA = document.getElementById("catA").value
    catB = document.getElementById("catB").value
    catC = document.getElementById("catC").value
    gas = document.getElementById("gas").value

    if (catA <= 0 && catB <= 0 && catC <= 0 && gas <= 0) {
        alert("Give a possitive number")
    } else {
        price_catA = Math.ceil(((catA - 2000) / 2000)) * 10 + baseCharge
        price_catB = Math.ceil(((catB - 2000) / 2000)) * 10 + baseCharge
        price_catC = Math.ceil(((catC - 2000) / 2000)) * 10 + baseCharge
        price_gas = Math.ceil(((gas - 2000) / 2000)) * 10 + baseCharge

        let final_price = price_catA + price_catB + price_catC + price_gas
        final_price *= exchangeRate

        document.getElementById("final-price").innerHTML = "€  " + final_price.toString()
    }


}
