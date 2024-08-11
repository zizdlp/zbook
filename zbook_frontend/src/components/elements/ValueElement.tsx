import ToolTip from "@/components/ToolTip";
import { ReactNode } from "react";
export default function ValueElement({
  tip,
  content,
}: {
  tip: string;
  content: ReactNode;
}) {
  return (
    <ToolTip message={tip}>
      <div className="border-[0.05rem] border-slate-300 dark:border-slate-700 rounded-md px-2 py-1 text-sm">
        {content}
      </div>
    </ToolTip>
  );
}
