'use client';

import { MouseEvent } from 'react';
import { useState, useEffect, useRef, useCallback } from 'react';
import ReactDOM from 'react-dom';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTimes } from '@fortawesome/free-solid-svg-icons';

import style from '@/styles/Modal.module.css';

import clsx from 'clsx';

export default function Modal({
  onClose,
  title,
  children,
}: {
  onClose: () => void;
  title: string;
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
        <div className={style.title}>
          <h3>{title}</h3>
          <button onClick={handleCloseClick}>
            <FontAwesomeIcon
              icon={faTimes}
              className="fa fa-clone mr-1 text-gray-500"
              size="lg"
            />
          </button>
        </div>
        {/* body */}
        <div>{children}</div>
      </div>
    </div>
  );

  return ReactDOM.createPortal(modalContent, document.body);
}
