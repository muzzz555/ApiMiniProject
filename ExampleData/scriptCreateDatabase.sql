create table type_user
(
    TypeNumber int          not null,
    TypeName   varchar(255) not null,
    constraint type_Customer_TypeName_uindex
        unique (TypeName),
    constraint type_Customer_TypeNumber_uindex
        unique (TypeNumber)
);

alter table type_user
    add primary key (TypeNumber);

create table company
(
    ID              int auto_increment,
    CompanyName     varchar(255)  not null,
    CompanyEmail    varchar(255)  not null,
    CompanyPhone    varchar(255)  not null,
    Address         varchar(255)  not null,
    Subdistrict     varchar(255)  not null,
    District        varchar(255)  not null,
    Province        varchar(255)  not null,
    Postcode        varchar(255)  not null,
    Password        varchar(255)  not null,
    Profile_Company varchar(255)  null,
    TypeNumber_User int default 2 not null,
    constraint Account_Company_ID_uindex
        unique (ID),
    constraint company_CompanyEmail_uindex
        unique (CompanyEmail),
    constraint company_CompanyPhone_uindex
        unique (CompanyPhone),
    constraint company_type_user_TypeNumber_fk
        foreign key (TypeNumber_User) references type_user (TypeNumber)
);

alter table company
    add primary key (ID);

create table type_work
(
    Type_Work_ID   int auto_increment,
    Type_Work_Name varchar(255) not null,
    constraint type_work_type_work_id_uindex
        unique (Type_Work_ID),
    constraint type_work_type_work_name_uindex
        unique (Type_Work_Name)
);

alter table type_work
    add primary key (Type_Work_ID);

create table user
(
    ID              int auto_increment,
    FirstName       varchar(255)  not null,
    LastName        varchar(255)  not null,
    Email           varchar(255)  not null,
    Password        varchar(255)  not null,
    Phone           varchar(255)  not null,
    TypeNumber_User int default 1 null,
    Profile_user    varchar(255)  null,
    Line            varchar(255)  null,
    Facebook        varchar(255)  null,
    Instagram       varchar(255)  null,
    constraint user_ID_uindex
        unique (ID),
    constraint user_Phone_uindex
        unique (Phone),
    constraint user_account_Email_uindex
        unique (Email),
    constraint user_type_user_TypeNumber_fk
        foreign key (TypeNumber_User) references type_user (TypeNumber)
);

alter table user
    add primary key (ID);

create table work_post_company
(
    Work_Post_ID            int auto_increment,
    Company_ID              int          not null,
    Type_Work_Number        int          not null,
    Name_Work               varchar(255) not null,
    Detail_Work             varchar(255) not null,
    Position                varchar(255) not null,
    Num_Person              int          not null,
    Price_Work_Min          varchar(255) not null,
    Price_Work_Max          varchar(255) not null,
    Education               varchar(255) not null,
    Image_Work_Post_Company varchar(255) not null,
    constraint work_post_company_work_post_id_uindex
        unique (Work_Post_ID),
    constraint work_post_company_type_work_type_work_id_fk
        foreign key (Type_Work_Number) references type_work (Type_Work_ID)
);

alter table work_post_company
    add primary key (Work_Post_ID);

create table work_post_freelance
(
    Work_Post_ID              int auto_increment,
    Type_Work_Number          int          not null,
    User_ID                   int          not null,
    Detail_Work               varchar(255) not null,
    Price_Post_Work           varchar(255) not null,
    Name_Work                 varchar(255) not null,
    Image_Work_Post_Freelance varchar(255) not null,
    constraint work_PostFreelance_work_Post_id_uindex
        unique (Work_Post_ID),
    constraint work_postfreelance_type_work_type_work_id_fk
        foreign key (Type_Work_Number) references type_work (Type_Work_ID),
    constraint work_postfreelance_user_ID_fk
        foreign key (User_ID) references user (ID)
);