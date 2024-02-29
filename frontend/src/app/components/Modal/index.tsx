import Image from "next/image";

import { useTheme } from "../../contexts/Theme";

type Props = {
  children: React.ReactNode;
  onClose: () => void;
};

const Modal = ({ children, onClose }: Props) => {
  const { theme, bgColor } = useTheme();

  const src = theme === "dark" ? "/light-close.png" : "/dark-close.png";

  return (
    <div className="fixed inset-0 z-50 flex flex-col items-center justify-center bg-black bg-opacity-50">
      <div className="flex flex-col">
        <button
          className="p-2 ml-auto rounded-full hover:bg-red-500"
          onClick={onClose}
        >
          <Image width={24} height={24} src={src} alt="Close" />
        </button>
        <div className={`p-8 mx-10 mb-10 ${bgColor} border rounded shadow-lg`}>
          {children}
        </div>
      </div>
    </div>
  );
};

export default Modal;
