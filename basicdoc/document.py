class Info:
    def __init__(self):
        self.Title = ""
        self.Description = ""
        self.TermsOfService = ""
        self.Version = ""
        self.License = None
        self.Contact = None

class Contact:
    def __init__(self):
        self.Name = ""
        self.Url = ""
        self.Email = ""

class License:
    def __init__(self):
        self.Name = ""
        self.Url = ""

class BasicInfo:
    def __init__(self, openapi: str, info: Info):
        self.openapi = openapi
        self.info = info

class Document:
    def __init__(self, info: BasicInfo):
        self.info = info

    # info: BasicInfo

    def a():
        print("test")