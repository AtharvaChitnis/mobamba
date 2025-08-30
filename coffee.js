'use strict';

function makeCoffee(x) {
  function boilWater() {
    console.log('boiling water');
    if (x === 95) {
      console.log('water boiled');
      function addCoffee() {
        console.log('adding coffee');
        function pourInCup() {
          console.log('sloWlly pouring in cup');
          console.log('coffee is ready');
        }
        pourInCup();
      }
      addCoffee();
    } else {
      console.log('boiling water again');
    }
  }
    boilWater();
}

module.exports = makeCoffee;
