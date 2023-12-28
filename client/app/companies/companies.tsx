'use client';

import { useState } from 'react';

import Searchbar from '@/ui/Searchbar';
import Button from '@/ui/Button';
import Modal from '@/ui/Modal';

import CompanyCard from '@/ui/companies/CompanyCard';
import CompanyForm from '@/ui/companies/CompanyForm';

import { Company } from '@/lib/types';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlus } from '@fortawesome/free-solid-svg-icons';

import style from '@/styles/company/CompanyPage.module.css';

export default function Companies({ companies }: { companies: Company[] }) {
  const [showModal, setShowModal] = useState(false);

  const onClose = () => {
    setShowModal(false);
  };

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
          <Modal title="Add Company" onClose={onClose}>
            <CompanyForm onClose={onClose} />
          </Modal>
        )
      }
      <Searchbar type="text" placeholder="Search for company" />
      <div className={style.cards}>
        {companies.map((company: Company) => (
          <CompanyCard key={company.id} company={company} />
        ))}
      </div>
    </div>
  );
}
