// ----------------------------------------------------------------
// Here we define our vendor dependencies
// ----------------------------------------------------------------
const vendors = {}
vendors["react"] = window.React;
vendors["react-dom"] = window.ReactDOM;

// ----------------------------------------------------------------
// Simulate require
// ----------------------------------------------------------------
window.require = (moduleName) => {
  if (!vendors[moduleName]) {
    throw new Error("Could not resolve module" + moduleName);
  }
  return vendors[moduleName];
}