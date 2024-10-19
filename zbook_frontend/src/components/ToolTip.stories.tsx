import type { Meta, StoryObj } from "@storybook/react";
import ToolTip from "./ToolTip"; // 确保路径正确
import React from "react";

// 定义元信息
const meta: Meta<typeof ToolTip> = {
  title: "Example/ToolTip",
  component: ToolTip,
  tags: ["autodocs"],
};

export default meta;

// 定义故事类型
type Story = StoryObj<typeof meta>;

// 定义一个基本故事
export const Default: Story = {
  args: {
    message: "This is a tooltip message!",
    children: (
      <div className="border-[0.05rem] border-slate-300 dark:border-slate-700 rounded-md px-2 py-1 text-sm">
        这是具体内容
      </div>
    ),
  },
};

// 如果有暗色主题，你可以添加一个故事来模拟它
export const DarkTheme: Story = {
  args: {
    message: "This is a tooltip message!",
    children: (
      <div className="border-[0.05rem] border-slate-300 dark:border-slate-700 rounded-md px-2 py-1 text-sm">
        这是具体内容
      </div>
    ),
  },
  parameters: {
    // 在 Storybook 中设置暗色主题
    backgrounds: { default: "dark" },
  },
};
