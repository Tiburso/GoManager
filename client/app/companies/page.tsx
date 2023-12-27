'use client';

import { useState, useEffect } from 'react';

import Searchbar from '@/ui/Searchbar';
import Button from '@/ui/Button';
import Modal from '@/ui/Modal';

import CompanyCards from '@/ui/companies/CompanyCards';
import CompanyForm from '@/ui/companies/CompanyForm';

import { Company } from '@/lib/types';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlus } from '@fortawesome/free-solid-svg-icons';

import style from '@/styles/company/CompanyPage.module.css';

const CompanyData: Company[] = [
  {
    id: '1',
    name: 'Google',
    company_portal: 'https://www.google.com/',
  },
  {
    id: '2',
    name: 'Facebook',
    company_portal: 'https://www.facebook.com/',
  },
  {
    id: '3',
    name: 'Amazon',
    company_portal: 'https://www.amazon.com/',
  },
  {
    id: '4',
    name: 'Apple',
    company_portal: 'https://www.apple.com/',
  },
  {
    id: '5',
    name: 'Microsoft',
    company_portal: 'https://www.microsoft.com/',
  },
  {
    id: '6',
    name: 'Netflix',
    company_portal: 'https://www.netflix.com/',
  },
];

export default function CompanyPage() {
  const [companies, setCompanies] = useState(CompanyData);
  const [showModal, setShowModal] = useState(false);

  const onClose = () => {
    setShowModal(false);
  };

  const onSave = (c: Company) => {
    // Here it should call the endpoint to save the company

    // TEMPORARY: generate a new id for the company based on the returned id
    // this would be done in the backend
    c.id = (companies.length + 1).toString();

    setCompanies([...companies, c]);
    setShowModal(false);
  };

  const onDelete = (id: string) => {
    // Here it should call the endpoint to delete the company

    setCompanies(companies.filter((c) => c.id !== id));
  };

  useEffect(() => {
    setCompanies(CompanyData);
  }, []);

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
          <Modal title="Add Company" onClose={onClose} saveText="Submit">
            <CompanyForm onClose={onClose} onSave={onSave} />
          </Modal>
        )
      }
      <Searchbar type="text" placeholder="Search for company" />
      <CompanyCards companies={companies} />
    </div>
  );
}
