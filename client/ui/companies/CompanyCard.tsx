"use client";

import { useState } from 'react';

import { Company } from '@/lib/types';

import Button from '@/ui/Button';
import Modal from '@/ui/Modal';
import CompanyForm from '@/ui/companies/CompanyForm';
import Link from 'next/link';

import style from '@/styles//company/CompanyCard.module.css';

import { updateCompany, deleteCompany } from '@/lib/companies';

export default function CompanyCard({ company }: { company: Company }) {
  const [showModal, setShowModal] = useState(false);

  const onClose = () => {
    setShowModal(false);
  }

  return (
    <div className={style.card}>
      <Link href={company.candidate_portal}>
        <h2>{company.name}</h2>
        <p>{company.candidate_portal}</p>
      </Link>

      <div className={style.btns}>
        {/* create an empty form here with the action to delete */}
        <form action={async () => await deleteCompany(company.id!)}>
          <Button className={style.delete} type="submit">Delete</Button>
        </form>

        <Button onClick={() => setShowModal(true)} type="submit">
          Edit
        </Button>
      </div>
      
      {showModal && (
        <Modal onClose={() => setShowModal(false)} title="Update company">
          <CompanyForm action={
            async (formData: FormData) => {
              try {
                // Here company id should never be null
                await updateCompany(company.id!, formData);
                onClose();
                return 'Company updated'
              } catch (error) {
                console.error(error);
                return 'Error updating company'
              }
            }
          } onClose={onClose} company={company} />
        </Modal>
      )}
    </div>
  );
}
