#!/usr/bin/env node

const toSpongeCase = (text) => {
    let result = '';
    let upper = false; // Start with lowercase

    for (let char of text) {
        if (char.match(/[a-zA-Z]/)) { // Check if the character is a letter
            result += upper ? char.toUpperCase() : char.toLowerCase();
            upper = !upper; // Toggle the case for the next letter
        } else {
            result += char; // Keep non-letter characters as they are
        }
    }

    return result;
}

// Get input from command line arguments
const inputText = process.argv.slice(2).join(' ');

if (inputText) {
    const spongeCasedText = toSpongeCase(inputText);
    console.log(spongeCasedText); // Output the SpongeCase text
} else {
    console.log("Please provide some text as input.");
}

