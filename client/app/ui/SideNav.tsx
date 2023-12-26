'use client';

// FontAwesome
import { config } from '@fortawesome/fontawesome-svg-core';
import '@fortawesome/fontawesome-svg-core/styles.css';
config.autoAddCss = false;

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faHome,
  faBriefcase,
  faBuilding,
} from '@fortawesome/free-solid-svg-icons';

import { usePathname } from 'next/navigation';
import Link from 'next/link';
import Image from 'next/image';

import clsx from 'clsx';
import styles from '../styles/SideNav.module.css';

export default function SideNav() {
  const pathname = usePathname();

  return (
    <div id="sidenav" className={styles.sidenav}>
      {/* add application icon this can be a temporary one from fontawesome */}
      <Image
        src="/logo_transparent.png"
        width={200}
        height={200}
        className="hidden md:block"
        alt="logo"
      />

      <Link href="/" className={clsx({ [styles.active]: pathname === '/' })}>
        <FontAwesomeIcon icon={faHome} />
        <span>Home</span>
      </Link>
      <Link
        href="/applications"
        className={clsx({ [styles.active]: pathname === '/applications' })}
      >
        <FontAwesomeIcon icon={faBriefcase} />
        <span>Applications</span>
      </Link>
      <Link
        href="/companies"
        className={clsx({ [styles.active]: pathname === '/companies' })}
      >
        <FontAwesomeIcon icon={faBuilding} />
        <span>Companies</span>
      </Link>
    </div>
  );
}
