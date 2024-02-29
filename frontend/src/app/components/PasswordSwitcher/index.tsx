import Image from "next/image";

import { useTheme } from "../../contexts/Theme";

type Props = {
  show: boolean;
  onClick: () => void;
};

const PasswordSwitcher = ({ show, onClick }: Props) => {
  const { theme, hoverColor } = useTheme();

  const { srcShow, srcHide } =
    theme === "dark"
      ? {
          srcShow: "/light-show-password.png",
          srcHide: "/light-hide-password.png",
        }
      : {
          srcShow: "/dark-show-password.png",
          srcHide: "/dark-hide-password.png",
        };

  return (
    <button className={`p-1 rounded-full ${hoverColor}`} onClick={onClick}>
      {show ? (
        <Image width={24} height={24} src={srcShow} alt="Show password" />
      ) : (
        <Image width={24} height={24} src={srcHide} alt="Hide password" />
      )}
    </button>
  );
};

export default PasswordSwitcher;
