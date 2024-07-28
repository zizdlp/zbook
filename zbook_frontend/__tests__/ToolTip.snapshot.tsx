import { render } from "@testing-library/react";
import ToolTip from "@/components/ToolTip";

it("renders ToolTip unchanged", () => {
  const message = "This is a tooltip";

  const { container } = render(
    <ToolTip message={message}>
      <button>Hover over me</button>
    </ToolTip>
  );
  expect(container).toMatchSnapshot();
});
