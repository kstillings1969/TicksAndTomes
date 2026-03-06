.eslintrc.cjs
{
  "root": true,
  "env": { "browser": true, "es2020": true },
  "extends": [
    "react-app",
    "react-app/jest"
  ],
  "rules": {
    "no-unused-vars": ["warn", { "argsIgnorePattern": "^_" }],
    "no-console": ["warn", { "allow": ["warn", "error"] }]
  }
}
