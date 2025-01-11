# Band Manager API Documentation

## Spis treści
- [Autentykacja](#autentykacja)
- [Grupy](#grupy)
- [Podgrupy](#podgrupy)
- [Wydarzenia](#wydarzenia)
- [Utwory](#utwory)
- [Ogłoszenia](#ogloszenia)

## Autentykacja

### Rejestracja użytkownika
- **URL**: `/api/verify/register`
- **Metoda**: `POST`
- **Body**:
```json
{
    "first_name": "string",
    "last_name": "string",
    "email": "string",
    "password": "string"
}
```
- **Odpowiedź**: 
  - Sukces (200): `{"message": "Registration successful"}`
  - Błąd (400/500): Komunikat błędu

### Logowanie
- **URL**: `/api/verify/login`
- **Metoda**: `POST`
- **Body**:
```json
{
    "email": "string",
    "password": "string"
}
```
- **Odpowiedź**:
```json
{
    "id": "uint",
    "first_name": "string",
    "last_name": "string",
    "email": "string",
    "groups": [
        {
            "id": "uint",
            "name": "string",
            "role": "string"
        }
    ]
}
```

## Grupy

### Tworzenie grupy
- **URL**: `/api/group/create`
- **Metoda**: `POST`
- **Body**:
```json
{
    "name": "string",
    "description": "string",
    "user_id": "uint"
}
```
- **Odpowiedź**:
```json
{
    "name": "string",
    "id": "uint",
    "role": "string"
}
```

### Dołączanie do grupy
- **URL**: `/api/group/join`
- **Metoda**: `POST`
- **Body**:
```json
{
    "user_id": "uint",
    "access_token": "string"
}
```
- **Odpowiedź**:
```json
{
    "user_role": "string",
    "user_group_id": "uint"
}
```

### Informacje o grupie
- **URL**: `/api/group/{group_id}/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**:
```json
{
    "name": "string",
    "description": "string",
    "access_token": "string" // Tylko dla managerów
}
```

### Lista grup użytkownika
- **URL**: `/api/group/user/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**:
```json
{
    "groups": [
        {
            "id": "uint",
            "name": "string",
            "description": "string",
            "role": "string",
            "members_count": "int"
        }
    ]
}
```

### Lista członków grupy
- **URL**: `/api/group/members/{group_id}/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**:
```json
{
    "members": [
        {
            "id": "uint",
            "first_name": "string",
            "last_name": "string",
            "email": "string",
            "role": "string"
        }
    ]
}
```

### Usunięcie członka
- **URL**: `/api/group/remove/{group_id}/{requester_id}/{user_id}`
- **Metoda**: `DELETE`
- **Odpowiedź**: 
  - Sukces (200): `{"message": "Member removed successfully"}`
  - Błąd (400/500): Komunikat błędu

### Aktualizacja roli członka
- **URL**: `/api/group/role/{group_id}/{user_id}/{requester_id}`
- **Metoda**: `PUT`
- **Body**:
```json
{
    "new_role": "string" // "manager", "moderator", lub "member"
}
```
- **Odpowiedź**: 
  - Sukces (200): `{"message": "Role updated successfully"}`
  - Błąd (400/500): Komunikat błędu

## Podgrupy

### Tworzenie podgrupy
- **URL**: `/api/subgroup/create`
- **Metoda**: `POST`
- **Body**:
```json
{
    "group_id": "uint",
    "name": "string",
    "description": "string",
    "user_id": "uint"
}
```
- **Odpowiedź**: Utworzony obiekt podgrupy

### Lista podgrup w grupie
- **URL**: `/api/subgroup/group/{group_id}/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**:
```json
{
    "subgroups": [
        {
            "id": "uint",
            "group_id": "uint",
            "name": "string",
            "description": "string"
        }
    ]
}
```
- **Kody błędów**:
  - 400: Nieprawidłowe ID
  - 401: Użytkownik nie należy do grupy
  - 500: Błąd serwera

### Informacje o podgrupie
- **URL**: `/api/subgroup/info/{subgroup_id}/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**: Obiekt podgrupy z szczegółami

### Aktualizacja podgrupy
- **URL**: `/api/subgroup/update/{subgroup_id}/{user_id}`
- **Metoda**: `PUT`
- **Body**:
```json
{
    "name": "string",
    "description": "string"
}
```
- **Odpowiedź**: 
  - Sukces (200): `{"message": "Subgroup updated successfully"}`
  - Błąd (400/500): Komunikat błędu

### Usunięcie podgrupy
- **URL**: `/api/subgroup/delete/{subgroup_id}/{user_id}`
- **Metoda**: `DELETE`
- **Odpowiedź**: 
  - Sukces (200): `{"message": "Subgroup deleted successfully"}`
  - Błąd (400/500): Komunikat błędu

### Dodawanie członków do podgrupy
- **URL**: `/api/subgroup/members/add/{subgroup_id}/{user_id}`
- **Metoda**: `POST`
- **Body**:
```json
{
    "user_ids": ["uint"]
}
```
- **Odpowiedź**: 
  - Sukces (200): `{"message": "Members added successfully"}`
  - Błąd (400/500): Komunikat błędu

### Usunięcie członka z podgrupy
- **URL**: `/api/subgroup/members/remove/{subgroup_id}/{member_id}/{requesting_user_id}`
- **Metoda**: `DELETE`
- **Odpowiedź**: 
  - Sukces (200): `{"message": "Member removed successfully"}`
  - Błąd (400/500): Komunikat błędu

## Wydarzenia

### Tworzenie wydarzenia
- **URL**: `/api/event/create`
- **Metoda**: `POST`
- **Body**:
```json
{
    "title": "string",
    "description": "string",
    "location": "string",
    "date": "timestamp",
    "group_id": "uint",
    "track_ids": ["uint"],
    "user_ids": ["uint"],
    "user_id": "uint"
}
```
- **Odpowiedź**: Utworzony obiekt wydarzenia

### Informacje o wydarzeniu
- **URL**: `/api/event/info/{event_id}/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**: Obiekt wydarzenia z szczegółami

### Aktualizacja wydarzenia
- **URL**: `/api/event/update/{event_id}/{user_id}`
- **Metoda**: `PUT`
- **Body**:
```json
{
    "title": "string",
    "description": "string",
    "location": "string",
    "date": "timestamp",
    "track_ids": ["uint"],
    "user_ids": ["uint"]
}
```
- **Odpowiedź**: 
  - Sukces (200): `{"message": "Event updated successfully"}`
  - Błąd (400/500): Komunikat błędu

### Lista wydarzeń grupy
- **URL**: `/api/event/group/{group_id}/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**:
```json
{
    "events": [Event objects]
}
```

### Lista utworów wydarzenia
- **URL**: `/api/event/tracks/{event_id}/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**:
```json
{
    "tracks": [Track objects]
}
```

### Usunięcie wydarzenia
- **URL**: `/api/event/delete/{event_id}/{user_id}`
- **Metoda**: `DELETE`
- **Odpowiedź**: 
  - Sukces (200): `{"message": "Event deleted successfully"}`
  - Błąd (400/500): Komunikat błędu

### Lista wydarzeń użytkownika
- **URL**: `/api/event/user/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**:
```json
{
    "events": [Event objects]
}
```

## Utwory

### Tworzenie utworu
- **URL**: `/api/track/create`
- **Metoda**: `POST`
- **Body**:
```json
{
    "title": "string",
    "description": "string",
    "group_id": "uint",
    "user_id": "uint"
}
```
- **Odpowiedź**: Utworzony obiekt utworu

### Dodawanie nut
- **URL**: `/api/track/notesheet`
- **Metoda**: `POST`
- **Body**:
```json
{
    "track_id": "uint",
    "user_id": "uint",
    "filepath": "string",
    "instrument": "string",
    "subgroup_ids": ["uint"]
}
```
- **Odpowiedź**: Utworzony obiekt nut

### Lista nut użytkownika
- **URL**: `/api/track/user/notesheets/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**:
```json
{
    "notesheets": [Notesheet objects]
}
```

### Lista utworów grupy
- **URL**: `/api/track/group/{group_id}/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**:
```json
{
    "tracks": [Track objects]
}
```

## Ogłoszenia

### Tworzenie ogłoszenia
- **URL**: `/api/announcement/create`
- **Metoda**: `POST`
- **Body**:
```json
{
    "title": "string",
    "description": "string",
    "priority": "uint",
    "group_id": "uint",
    "sender_id": "uint",
    "subgroup_ids": ["uint"]
}
```
- **Odpowiedź**: Utworzony obiekt ogłoszenia

### Usunięcie ogłoszenia
- **URL**: `/api/announcement/delete/{announcement_id}/{user_id}`
- **Metoda**: `DELETE`
- **Odpowiedź**: 
  - Sukces (200): `{"message": "Announcement deleted successfully"}`
  - Błąd (400/500): Komunikat błędu

### Lista ogłoszeń użytkownika
- **URL**: `/api/announcement/user/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**:
```json
{
    "announcements": [Announcement objects]
}
```

### Lista ogłoszeń grupy
- **URL**: `/api/announcement/group/{group_id}/{user_id}`
- **Metoda**: `GET`
- **Odpowiedź**:
```json
{
    "announcements": [Announcement objects]
}
```