import Image from "next/image";

import { useTheme } from "../../contexts/Theme";

type Props = {
  children: React.ReactNode;
  onClose: () => void;
};

const Modal = ({ children, onClose }: Props) => {
  const { lights } = useTheme();

  const { divTheme, srcTheme } = lights
    ? {
        divTheme: "bg-white border-black",
        srcTheme: "/dark-close.png",
      }
    : {
        divTheme: "bg-black border-white",
        srcTheme: "/light-close.png",
      };

  return (
    <div className="fixed inset-0 z-50 flex flex-col items-center justify-center bg-black bg-opacity-50">
      <div className="flex flex-col">
        <button
          className="p-2 ml-auto rounded-full hover:bg-red-500"
          onClick={onClose}
        >
          <Image width={24} height={24} src={srcTheme} alt="Close" />
        </button>
        <div className={`p-8 mx-10 mb-10 border rounded shadow-lg ${divTheme}`}>
          {children}
        </div>
      </div>
    </div>
  );
};

export default Modal;
