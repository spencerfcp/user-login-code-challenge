import { validateStringField } from "./validate";
import { test, expect } from "@playwright/test";

test.describe("form validation", () => {
  test.describe("username validation", () => {
    test("it should fail if invalid length", () => {
      expect(validateStringField("username", "J")).toBeFalsy();
      expect(validateStringField("username", "j".repeat(81))).toBeFalsy();
    });
    test("it should pass with enough characters", () => {
      expect(validateStringField("username", "JeffSpencer")).toBeTruthy();
    });
  });

  test.describe("password validation", () => {
    test("it should fail if password invalid", () => {
      expect(validateStringField("password", "J")).toBeFalsy();
      expect(validateStringField("password", "password")).toBeFalsy();
      expect(
        validateStringField("password", "stillweak!passsword")
      ).toBeFalsy();
    });
    test("it should pass if password is valid", () => {
      expect(
        validateStringField("password", "TheBestPassword1!!")
      ).toBeTruthy();
    });
  });
});
