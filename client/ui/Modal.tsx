'use client';

import { MouseEvent } from 'react';
import { useState, useEffect, useRef, useCallback } from 'react';
import ReactDOM from 'react-dom';

import style from '@/styles/Modal.module.css';

import clsx from 'clsx';

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

  useEffect(() => {
    setMounted(true);
    document.addEventListener('click', handleOutsideClick, true);
    return () => {
      document.removeEventListener('click', handleOutsideClick, true);
    };
  }, [handleOutsideClick]);

  const modalContent = (
    <div className={clsx(style.modal, { [style.active]: mounted })}>
      <div ref={ref} className={style.wrapper}>
        <div className="border-b border-gray-200 px-4 py-2">
          <button onClick={handleCloseClick}>x</button>
        </div>
        <div className="px-4 py-2">{children}</div>
      </div>
    </div>
  );

  return ReactDOM.createPortal(modalContent, document.body);
}
