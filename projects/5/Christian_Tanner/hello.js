var Module = (function() {
  var Module = {};

  // The function to load the WASM module
  Module.onRuntimeInitialized = function() {
    // Function to call the WebAssembly function
    var generateGreeting = Module.cwrap('generate_greeting', 'string', ['string', 'number', 'number']);
    var freeBuffer = Module.cwrap('free_buffer', null, ['number']);

    // Event listener for the form submission
    document.getElementById('submit').addEventListener('click', function() {
      // Get user input
      var name = document.getElementById('name').value;
      var age = parseInt(document.getElementById('age').value, 10);
      var repeat = parseInt(document.getElementById('repeat').value, 10);

      // Call the WebAssembly function to get the greeting
      var resultPtr = generateGreeting(name, age, repeat);

      // Convert result pointer to a string
      var result = Module.UTF8ToString(resultPtr);

      // Display the greeting in the DOM
      document.getElementById('output').textContent = result;

      // Free the memory allocated for the result
      freeBuffer(resultPtr);
    });
  };

  return Module;
})();
