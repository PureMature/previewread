// Package starbox provides a comprehensive set of utilities for building and managing Starlark virtual machines with ease.
//
// # Module Sources
//
// Starbox supports loading modules from various sources, including built-in modules from Starlet, custom modules added by the user, and dynamic modules resolved by name on demand.
//
// Built-in Modules
//
// Use SetModuleSet(modSet ModuleSetName) to select a predefined set of modules to preload before execution. Available sets include:
//   - EmptyModuleSet: No modules.
//   - SafeModuleSet: Safe modules without access to the file system or network.
//   - NetworkModuleSet: Safe modules plus network modules.
//   - FullModuleSet: All available modules.
//
// Adding Custom Modules
//
// - AddModuleLoader(moduleName string, moduleLoader starlet.ModuleLoader): Adds a custom module loader. Members can be accessed in the script via load("module_name", "member_name") or member_name.
// - AddModuleFunctions(name string, funcs FuncMap): Adds a module with custom functions. Functions can be accessed in the script via load("module_name", "func_name") or module_name.func_name.
// - AddModuleData(moduleName string, moduleData starlark.StringDict): Adds a module with custom data. Data can be accessed in the script via load("module_name", "key") or module_name.key.
// - AddStructFunctions(name string, funcs FuncMap): Adds a module with custom struct functions. Functions can be accessed in the script via load("struct_name", "func_name") or struct_name.func_name.
// - AddStructData(structName string, structData starlark.StringDict): Adds a module with custom struct data. Data can be accessed in the script via load("struct_name", "key") or struct_name.key.
//
// Dynamic Module Loading
//
// - SetDynamicModuleLoader(loader DynamicModuleLoader): Sets a dynamic module loader, which loads modules based on their names before execution.
//
// Adding Named Modules
//
// - AddNamedModules(moduleNames ...string): Adds built-in or custom modules by their names.
// - AddModulesByName(moduleNames ...string): Alias for AddNamedModules.
//
// # Module Loading Priority
//
// Modules are loaded in the following order of priority before execution:
// 1. Preloaded Starlet modules from predefined sets and named Starlet modules.
// 2. Custom modules added by users, overriding any preloaded Starlet modules with the same names.
// 3. Dynamically loaded modules based on their names just before execution.
// 4. If a module name is not found in any of the preloaded, custom, or dynamic modules, an error is returned.
package previewread
