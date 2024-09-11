import React from "react";
import {
  FaGithub,
  FaDiscord,
  FaGoogle,
  FaFacebook,
  FaYoutube,
  FaTwitter,
  FaTwitch,
  FaGitlab,
  FaInstagram,
  FaLine,
} from "react-icons/fa"; // Import any other icons as needed
import { FaBilibili } from "react-icons/fa6";
import { BiLogoGmail } from "react-icons/bi";

import { IoBook, IoShareSocialSharp } from "react-icons/io5";
import { IconType } from "react-icons/lib";
import { MdCloudSync, MdOutlineVisibility } from "react-icons/md";

// Define a map of icon names to corresponding React icons
const iconMap: Record<string, IconType> = {
  github: FaGithub,
  discord: FaDiscord,
  google: FaGoogle,
  facebook: FaFacebook,
  youtube: FaYoutube,
  twitter: FaTwitter,
  twitch: FaTwitch,
  gitlab: FaGitlab,
  instagram: FaInstagram,
  bilibili: FaBilibili,
  line: FaLine,
  gmail: BiLogoGmail,
  MdCloudSync: MdCloudSync,
  MdOutlineVisibility: MdOutlineVisibility,
  IoBook: IoBook,
  // Add more icon mappings here
};

// Default icon to use if the provided iconName is not found in the map
const DefaultIcon = IoShareSocialSharp;

export default function IconItem({
  iconName,
  className,
}: {
  iconName: string;
  className?: string;
}) {
  // Determine which icon to use based on the iconName prop
  const IconComponent: IconType = iconMap[iconName] || DefaultIcon;

  return <IconComponent className={className} />;
}
