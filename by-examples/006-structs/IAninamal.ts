interface IAnimal {
    name : string;
    sound: string;
}

class Animal implements IAnimal{
    name = "generic animal"
    sound = "generic sound"
    speak(): string {
        return `The ${this.name} does ${this.sound}`
    }
}

class Dog extends Animal {
    name = "dog"
    sound = "woff woff"
}

class Cat extends Animal {
    name = "cat"
    sound = "meaw"
}

const a = 1
var animal : Animal;

if( a ==1 ){
    animal = new Dog()
} else {
    animal = new Cat()
}

console.log(animal.speak())