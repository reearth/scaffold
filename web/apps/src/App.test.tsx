import { render, screen } from "@testing-library/react";
import { expect, test } from "vitest";
import "@testing-library/jest-dom";

import App from "./App";

test("renders learn react link", () => {
  render(<App />);
  const linkElement = screen.getByText("テスト用のページです");
  expect(linkElement).toBeVisible();
});
