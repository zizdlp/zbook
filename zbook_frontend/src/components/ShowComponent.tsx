import { ReactNode } from "react";
export default function ShowComponent({
  show,
  children,
}: {
  show: boolean;
  children: ReactNode;
}) {
  if (show) {
    return <>{children}</>;
  } else {
    return <></>;
  }
}
