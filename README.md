# Copy Value
### Copy between 2 pointers under the following rule
- For basic types (bool, int, float, string), simply the value of it are copied.
- For slice or array of same type, each element is copied until reaching the nearest index bound.
- For map of same key-value type, all key-value pairs in the source map are copied to the destination map.
- For struct, copy is attempted between fields of same name.
	- The rule for each field is same as the one forementioned.
	- Types of source and destionation struct's need not be same.
