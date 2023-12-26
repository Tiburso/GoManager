'use client';

import { MouseEvent } from 'react';
import { useState, useEffect, useRef, useCallback } from 'react';
import ReactDOM from 'react-dom';

export default function Modal({
  onClose,
  children,
}: {
  onClose: () => void;
  children: React.ReactNode;
}) {
  const [mounted, setMounted] = useState(false);
  const ref = useRef<HTMLDivElement>(null);

  const handleCloseClick = (e: MouseEvent) => {
    e.preventDefault();
    onClose();
  };

  const handleOutsideClick = useCallback(
    (e: any) => {
      if (!ref?.current?.contains(e.target as Node)) {
        onClose();
      }
    },
    [onClose],
  );

  useEffect(() => setMounted(true), []);

  useEffect(() => {
    document.addEventListener('click', handleOutsideClick, true);
    return () => {
      document.removeEventListener('click', handleOutsideClick, true);
    };
  }, [handleOutsideClick]);

  const modalContent = (
    <div className="fixed top-0 left-0 z-[1040] bg-black w-screen h-screen opacity-50 transition-all duration-300 ease-in-out">
      <div ref={ref}>
        <div className="mx-auto my-10 max-w-md w-full bg-white rounded-xl shadow-md">
          <div className="border-b border-gray-200 px-4 py-2">
            <a href="#" onClick={handleCloseClick}>
              x
            </a>
          </div>
          <div className="px-4 py-2">{children}</div>
        </div>
      </div>
    </div>
  );

  return mounted ? ReactDOM.createPortal(modalContent, document.body) : null;
}
