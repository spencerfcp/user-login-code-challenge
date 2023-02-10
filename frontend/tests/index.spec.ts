import { test, expect } from "@playwright/test";

test.describe("index/login", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto("http://localhost:3000/");
  });

  test("user login", async ({ page }) => {
    await page.route("**/login", async (route) => {
      const json = {
        username: "ValidUsername",
      };
      await route.fulfill({
        status: 200,
        body: JSON.stringify(json),
      });
    });

    const successMessage = page.getByTestId("login_successful_message");
    await expect(successMessage).toBeHidden();
    const button = page.getByTestId("login_submit_button");

    const input = page.getByTestId("login_username");
    await input.fill("");
    await page.getByTestId("login_password").fill("TheBestPassword1!!!");

    await input.fill("ValidUsername");

    await button.click();
    await expect(successMessage).toBeVisible();
  });

  test("invalid - user login", async ({ page }) => {
    await page.route("**/login", async (route) => {
      const json = {
        invalidCredentials: true,
      };
      await route.fulfill({
        status: 200,
        body: JSON.stringify(json),
      });
    });

    const errorMessage = page.getByTestId("login_invalid_creds_error");
    await expect(errorMessage).toBeHidden();
    const button = page.getByTestId("login_submit_button");

    const input = page.getByTestId("login_username");
    await input.fill("");
    await page.getByTestId("login_password").fill("TheBestPassword1!!!");

    await input.fill("ValidUsername");

    await button.click();
    await expect(errorMessage).toBeVisible();
  });
});
