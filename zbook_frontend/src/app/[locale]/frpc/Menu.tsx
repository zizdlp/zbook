import React from "react";
import MenuItem from "./MenuItem";
import { FaBookOpen, FaDiscord, FaGithub } from "react-icons/fa";
export default function Menu() {
  const menuItems = [
    {
      href: "/introduction",
      text: "Documentation",
      selected: true,
      icon: FaBookOpen,
    },
    {
      href: "https://loopholelabs.io/discord",
      text: "discord",
      selected: false,
      icon: FaDiscord,
    },
    {
      href: "https://github.com/loopholelabs/frpc-go",
      text: "frpc",
      selected: false,
      icon: FaGithub,
    },
  ];

  return (
    <>
      {menuItems.map((item, index) => (
        <MenuItem key={index} {...item} />
      ))}
    </>
  );
}
