import React from "react";
import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";
import TimeElement from "@/components/TimeElement";

// Mock the useTranslations hook
jest.mock("next-intl", () => ({
  useTranslations:
    () => (key: keyof typeof translations, values?: { duration: number }) => {
      const translations = {
        JustNow: "Just now",
        MinuteAgo: `${values?.duration} minute(s) ago`,
        MinuteAfter: `In ${values?.duration} minute(s)`,
        HourAgo: `${values?.duration} hour(s) ago`,
        HourAfter: `In ${values?.duration} hour(s)`,
        DayAgo: `${values?.duration} day(s) ago`,
        DayAfter: `In ${values?.duration} day(s)`,
      };
      return translations[key];
    },
}));

describe("TimeElement", () => {
  it("renders the correct translation for time just now", () => {
    const now = new Date().toISOString();
    const { container } = render(<TimeElement timeInfo={now} />);
    expect(screen.getByText("Just now")).toBeInTheDocument();
    expect(container).toMatchSnapshot();
  });

  it("renders the correct translation for minutes ago", () => {
    const now = new Date();
    const minutesAgo = new Date(
      now.getTime() - 5 * 60 * 1000 - 10
    ).toISOString();
    const { container } = render(<TimeElement timeInfo={minutesAgo} />);
    expect(screen.getByText("5 minute(s) ago")).toBeInTheDocument();
    expect(container).toMatchSnapshot();
  });

  it("renders the correct translation for minutes after", () => {
    const now = new Date();
    const minutesAfter = new Date(
      now.getTime() + 5 * 60 * 1000 + 10
    ).toISOString();
    const { container } = render(<TimeElement timeInfo={minutesAfter} />);
    expect(screen.getByText("In 5 minute(s)")).toBeInTheDocument();
    expect(container).toMatchSnapshot();
  });

  it("renders the correct translation for hours ago", () => {
    const now = new Date();
    const hoursAgo = new Date(
      now.getTime() - 2 * 60 * 60 * 1000 - 10
    ).toISOString();
    const { container } = render(<TimeElement timeInfo={hoursAgo} />);
    expect(screen.getByText("2 hour(s) ago")).toBeInTheDocument();
    expect(container).toMatchSnapshot();
  });

  it("renders the correct translation for hours after", () => {
    const now = new Date();
    const hoursAfter = new Date(
      now.getTime() + 2 * 60 * 60 * 1000 + 10
    ).toISOString();
    const { container } = render(<TimeElement timeInfo={hoursAfter} />);
    expect(screen.getByText("In 2 hour(s)")).toBeInTheDocument();
    expect(container).toMatchSnapshot();
  });

  it("renders the correct translation for days ago", () => {
    const now = new Date();
    const daysAgo = new Date(
      now.getTime() - 3 * 24 * 60 * 60 * 1000 - 10
    ).toISOString();
    const { container } = render(<TimeElement timeInfo={daysAgo} />);
    expect(screen.getByText("3 day(s) ago")).toBeInTheDocument();
    expect(container).toMatchSnapshot();
  });

  it("renders the correct translation for days after", () => {
    const now = new Date();
    const daysAfter = new Date(
      now.getTime() + 3 * 24 * 60 * 60 * 1000 + 10
    ).toISOString();
    const { container } = render(<TimeElement timeInfo={daysAfter} />);
    expect(screen.getByText("In 3 day(s)")).toBeInTheDocument();
    expect(container).toMatchSnapshot();
  });
});
