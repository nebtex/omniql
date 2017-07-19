var Benchmark = require('benchmark');
 
var suite = new Benchmark.Suite;
p = {0:0}
j = {a:0, b:4}
// add tests
suite.add('RegExp#test', function() {
 p[0]=1;
 p[30]=1;

})
.add('String#indexOf', function() {
 j.a = 3;
 j.b = 4;
})
.add('String#match', function() {
 p[0]=1;
})
// add listeners
.on('cycle', function(event) {
  console.log(String(event.target));
})
.on('complete', function() {
  console.log('Fastest is ' + this.filter('fastest').map('name'));
})
// run async
.run({ 'async': true });