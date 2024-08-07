import { BsFillBookmarkCheckFill } from "react-icons/bs";
import { FaInfoCircle } from "react-icons/fa";
import { MdError, MdTipsAndUpdates } from "react-icons/md";
import { TiWarning } from "react-icons/ti";
type ParentType = "note" | "warning" | "info" | "tip" | "error";

export const iconTypes = {
  note: BsFillBookmarkCheckFill,
  warning: TiWarning,
  info: FaInfoCircle,
  tip: MdTipsAndUpdates,
  error: MdError,
};

const bgColors: Record<ParentType, string> = {
  note: "bg-cyan-600/75 dark:bg-cyan-500/25",
  warning: "bg-yellow-600/75 dark:bg-yellow-500/25",
  info: "bg-green-600/75 dark:bg-green-500/25",
  tip: "bg-lime-600/75 dark:bg-lime-500/25",
  error: "bg-red-600/75 dark:bg-red-500/25",
};

export const getParentTypeAndBg = (
  parent: HTMLElement | null
): { parentType: ParentType; bg: string } => {
  if (!parent) return { parentType: "note", bg: bgColors.note };
  const classList = parent.getAttribute("class");
  if (classList) {
    if (classList.includes("adm-note"))
      return { parentType: "note", bg: bgColors.note };
    if (classList.includes("adm-warning"))
      return { parentType: "warning", bg: bgColors.warning };
    if (classList.includes("adm-info"))
      return { parentType: "info", bg: bgColors.info };
    if (classList.includes("adm-tip"))
      return { parentType: "tip", bg: bgColors.tip };
    if (classList.includes("adm-error"))
      return { parentType: "error", bg: bgColors.error };
  }

  return { parentType: "note", bg: bgColors.note };
};
