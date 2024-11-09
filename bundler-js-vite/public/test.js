// Basic test function
export function test(description, callback) {
  try {
    callback();
    console.log(`✔️  ${description}`);
  } catch (error) {
    console.error(`❌  ${description}`);
    console.error(error);
  }
}

// Assertion function
export function assertEquals(actual, expected) {
  if (actual !== expected) {
    throw new Error(`Expected ${expected}, but got ${actual}`);
  }
}