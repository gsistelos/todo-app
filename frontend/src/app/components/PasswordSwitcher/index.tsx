import Image from "next/image";

import { useTheme } from "../../contexts/Theme";

type Props = {
  show: boolean;
  onClick: () => void;
};

const PasswordSwitcher = ({ show, onClick }: Props) => {
  const { lights } = useTheme();

  const { hover, showSrc, hideSrc } = lights
    ? {
        hover: "hover:bg-gray-200",
        showSrc: "/dark-show-password.png",
        hideSrc: "/dark-hide-password.png",
      }
    : {
        hover: "hover:bg-gray-800",
        showSrc: "/light-show-password.png",
        hideSrc: "/light-hide-password.png",
      };

  return (
    <button className={`p-1 rounded-full ${hover}`} onClick={onClick}>
      {show ? (
        <Image width={24} height={24} src={showSrc} alt="Show password" />
      ) : (
        <Image width={24} height={24} src={hideSrc} alt="Hide password" />
      )}
    </button>
  );
};

export default PasswordSwitcher;
