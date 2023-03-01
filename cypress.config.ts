import { defineConfig } from "cypress";

export default defineConfig({
  e2e: {
    baseUrl: "http://localhost:4200",
    video: false,
    screenshotOnRunFailure: false
  },

  component: {
    devServer: {
      framework: "angular",
      bundler: "webpack",
    },
    specPattern: "**/*.cy.ts",
    video: false,
    screenshotOnRunFailure: false
  },
});
