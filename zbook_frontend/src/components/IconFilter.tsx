import { FaDiscord, FaGithub } from "react-icons/fa";
import { IconType } from "react-icons/lib";

const IconText = ({
  Icon,
  class_name,
}: {
  Icon: IconType;
  class_name: string;
}) => <Icon className={class_name} />;
export default function IconFilter({
  icon_name,
  class_name,
}: {
  icon_name: string;
  class_name: string;
}) {
  if (icon_name == "discord") {
    return <IconText Icon={FaDiscord} class_name={class_name} />;
  } else if (icon_name == "github") {
    return <IconText Icon={FaGithub} class_name={class_name} />;
  }
  return <></>;
}
