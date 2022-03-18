# Qt 工程规范

## 编码与格式化

- 编码必须为 UTF-8，换行符必须为 LF（即 \n）
- C/C++ 代码必须用 clang-format 格式化
- QML 代码必须用 Qt Creator 格式化
- QML 中嵌入的 JavaScript 代码必须用 Visual Studio Code 格式化
- 无 Qt 依赖的 JavaScript 代码必须用 TypeScript 编写

## 命名规范

### 通用

- 源码文件命名必须为 camel_case 风格

### C/C++ & JavaScript/TypeScript

- 私有成员命名必须为 camelCase 风格
- 公有成员命名必须为 CamelCase 风格

### QML

- 组件类型+功能命名，比如 comboBoxUsers，代表用户下拉框
