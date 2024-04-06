'use client';

import { useEffect } from 'react';

import ReactDOM from 'react-dom';

type Props = {
  onClose?: () => void;
  children: React.ReactNode;
};

const Modal = ({ onClose, children }: Props) => {
  useEffect(() => {
    const handleEscape = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        onClose?.();
      }
    };

    document.addEventListener('keydown', handleEscape);
    document.body.style.overflow = 'hidden';

    return () => {
      document.removeEventListener('keydown', handleEscape);
      document.body.style.overflow = 'auto';
    };
  }, []);

  return (
    ReactDOM.createPortal(
      <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50" onClick={onClose}>
        <div className="bg-primary border border-contrast rounded-lg text-contrast" onClick={(event) => event.stopPropagation()}>
          {children}
        </div>
      </div>,
      document.getElementById('modal-root') as HTMLElement)
  );
};

export default Modal;
