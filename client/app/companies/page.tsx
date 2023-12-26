'use client';

import { useState } from 'react';

import Searchbar from '@/app/ui/Searchbar';
import Button from '@/app/ui/Button';
import Modal from '@/app/ui/Modal';
import CompanyCards from '@/app/ui/companies/CompanyCards';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlus } from '@fortawesome/free-solid-svg-icons';

import style from '@/app/styles/company/CompanyPage.module.css';

export default function CompanyPage() {
  const [showModal, setShowModal] = useState(false);

  return (
    <div className="w-4/5 mx-auto">
      <div className={`${style.title}`}>
        <h1 className="text-4xl text-left font-bold">Your Companies</h1>

        <Button
          onClick={() => {
            setShowModal(true);
          }}
        >
          <FontAwesomeIcon icon={faPlus} className="fa fa-clone fa-xs mr-1" />
          Add company
        </Button>
      </div>
      {
        /* Modal should be a component that is only visible when showModal is true */
        showModal && (
          <Modal onClose={() => setShowModal(false)}>
            <h1>Modal</h1>
          </Modal>
        )
      }
      <Searchbar type="text" placeholder="Search for company" />
      <CompanyCards />
    </div>
  );
}
