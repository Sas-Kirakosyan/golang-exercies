
//1 Create an object car with properties:
// brand: "Toyota"
// model: "Corolla"
// year: 2020
// Add a method displayInfo() that logs:
// "This is a Toyota Corolla from 2020."
// Call the method.


//2Create an object person with:
// name: "Alice"
// age: 25
// greet() method that logs:
// "Hello, my name is Alice and I am 25 years old."
// Use this to access properties.
// Call person.greet().

//3Create an object original with { a: 1, b: 2 }.
// Assign it to a new variable copy.
// Modify copy.a = 10.
// Log both original and copy. Explain why they changed.

//3Create obj1 = { x: 1, y: { z: 2 } }.
// Make a shallow copy (obj2)  spread ({...obj1}).
// Modify obj2.y.z = 100.
// Log both objects and explain the result.

const user = {
  id: 101,
  name: "John",
  age: 30,
  address: { city: "New York", country: "USA" }
};

//4  Extract name, age, and city using destructuring.
// Log them in one console.log:
// "John, 30, New York".


const product = { name: "Laptop", price: 999, stock: 5 };
//5 get all keys into array
// get all values into array
// get all key value payers like this
const arr = [
  [
    "name",
    "Laptop"
  ],
  [
    "price",
    999
  ],
  [
    "stock",
    5
  ]
]


//6
// Create objA = { a: 1, b: 2 } and objB = { b: 3, c: 4 }.
// Merge them into a new object objC where objB overwrites objA's properties.
// Log objC (should be { a: 1, b: 3, c: 4 }).


//7
let obj = {
  "John": 100,
  "Pete": 300,
  "Mary": 250
};

// find sum of numbers in obj


//8 
const arr = [
  {
    "name": "Charlie",
    "age": 32,
    "salary": 57945
  },
  {
    "name": "Alice",
    "age": 27,
    "salary": 31658
  },
  {
    "name": "Hannah",
    "age": 50,
    "salary": 73250
  },
  {
    "name": "George",
    "age": 53,
    "salary": 31184
  },
  {
    "name": "Ethan",
    "age": 48,
    "salary": 39918,
  }
]

// find average salary
// add bonus for each men which equal salary 15%