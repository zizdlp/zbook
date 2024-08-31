"use client";
import { MdDarkMode, MdLightMode } from "react-icons/md";
import { useTheme } from "next-themes";
import { useState, useEffect, useContext } from "react";
import { OperationContext } from "@/providers/OperationProvider";
import NavBarIcon from "./NavBarIcon";

export default function Theme() {
  const { theme, setTheme } = useTheme();
  const [mounted, setMounted] = useState(false);
  const { setMutationToggleTheme, mutationToggleTheme } =
    useContext(OperationContext);

  useEffect(() => {
    setMounted(true);
  }, []);

  return (
    <NavBarIcon
      Icon={theme === "dark" ? MdDarkMode : MdLightMode}
      onClick={() => {
        if (theme === "dark") {
          setTheme("light");
        } else {
          setTheme("dark");
        }
        setMutationToggleTheme(!mutationToggleTheme);
      }}
      mounted={mounted}
    />
  );
}
