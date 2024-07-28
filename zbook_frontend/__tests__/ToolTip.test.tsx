import React from "react";
import { render, screen, fireEvent } from "@testing-library/react";
import "@testing-library/jest-dom";
import ToolTip from "@/components/ToolTip";

describe("ToolTip", () => {
  it("shows the message when hovering over the child element", () => {
    const message = "This is a tooltip";
    render(
      <ToolTip message={message}>
        <button>Hover over me</button>
      </ToolTip>
    );

    // Initially, the tooltip should not be visible
    expect(screen.queryByText(message)).not.toBeInTheDocument();

    // Simulate mouse enter event
    fireEvent.mouseEnter(screen.getByText("Hover over me"));

    // Now the tooltip should be visible
    expect(screen.getByText(message)).toBeInTheDocument();

    // Simulate mouse leave event
    fireEvent.mouseLeave(screen.getByText("Hover over me"));

    // The tooltip should be hidden again
    expect(screen.queryByText(message)).not.toBeInTheDocument();
  });
});
