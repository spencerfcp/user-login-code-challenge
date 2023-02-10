import { test, expect } from "@playwright/test";

test.describe("signup", () => {
  test.beforeEach(async ({ page }) => {
    await page.goto("http://localhost:3000/signup");
  });

  test("successful - user sign up", async ({ page }) => {
    await page.route("**/user", async (route) => {
      const json = {
        user: {
          username: "validuser",
        },
      };
      await route.fulfill({
        status: 200,
        body: JSON.stringify(json),
      });
    });

    const successMessage = page.getByTestId("signup_successful_message");
    const button = page.getByTestId("signup_submit_button");
    await expect(successMessage).toBeHidden();
    await page.getByTestId("signup_username").fill("Whatever");
    await page.getByTestId("signup_password").fill("TheBestPassword1!!!");
    await expect(button).toBeEnabled();
    await button.click();
    await expect(successMessage).toBeVisible();
  });

  test("already exists - user sign up", async ({ page }) => {
    await page.route("**/user", async (route) => {
      const json = {
        usernameAlreadyExists: true,
      };
      await route.fulfill({
        status: 200,
        body: JSON.stringify(json),
      });
    });

    const errorMessage = page.getByTestId("signup_user_exists_error");
    await expect(errorMessage).toBeHidden();

    await page.getByTestId("signup_username").fill("Whatever");
    await page.getByTestId("signup_password").fill("TheBestPassword1!!!");
    await page.getByTestId("signup_submit_button").click();
    await expect(page.getByTestId("signup_user_exists_error")).toBeVisible();
  });
});
