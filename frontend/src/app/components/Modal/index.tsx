import Image from "next/image";

import { useTheme } from "../../contexts/Theme";

type Props = {
  children: React.ReactNode;
  onClose: () => void;
};

const Modal = ({ children, onClose }: Props) => {
  const { lights } = useTheme();

  const { container, hover, src } = lights
    ? {
        container: "border-black bg-white",
        hover: "hover:bg-gray-200",
        src: "/dark-close.png",
      }
    : {
        container: "border-white bg-black",
        hover: "hover:bg-gray-800",
        src: "/light-close.png",
      };

  return (
    <div className="fixed inset-0 z-50 flex flex-col items-center justify-center bg-black bg-opacity-50">
      <div className="flex flex-col">
        <button
          className={`ml-auto mb-2 p-2 rounded-full ${hover}`}
          onClick={onClose}
        >
          <Image src={src} alt="Close" width={24} height={24} />
        </button>
        <div className={`mx-10 p-8 border rounded shadow-lg ${container}`}>
          {children}
        </div>
      </div>
    </div>
  );
};

export default Modal;
