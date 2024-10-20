import type { Meta, StoryObj } from "@storybook/react";
import { fn } from "@storybook/test";

import { Editor } from "./index";

const meta = {
  title: "Editor",
  component: Editor,
  // This component will have an automatically generated Autodocs entry: https://storybook.js.org/docs/writing-docs/autodocs
  tags: ["autodocs"],
  parameters: {
    // More on how to position stories at: https://storybook.js.org/docs/configure/story-layout
    layout: "fullscreen",
  },
  args: {
    disableExtensions: [],
  },
} satisfies Meta<typeof Editor>;

export default meta;
type Story = StoryObj<typeof meta>;

export const DefaultEditor: Story = {
  args: {
    defaultValue: `# Welcome

    Just an easy to use **Markdown** editor with \`slash commands\``,
  },
};
