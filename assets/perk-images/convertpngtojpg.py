from PIL import Image
import os

def batch_convert_png_to_jpg(input_folder, output_folder, quality=85):
    if not os.path.exists(output_folder):
        os.makedirs(output_folder)

    for file_name in os.listdir(input_folder):
        if file_name.endswith(".png"):
            input_path = os.path.join(input_folder, file_name)
            output_path = os.path.join(output_folder, file_name.replace(".png", ".jpg"))
            png_to_jpg(input_path, output_path, quality)

def png_to_jpg(input_path, output_path, quality=85):
    image = Image.open(input_path)
    if image.mode in ("RGBA", "P"):
        image = image.convert("RGB")
    image.save(output_path, "JPEG", quality=quality)

# Example usage:
cwd = os.getcwd()
input_folder = os.path.join(cwd, "profileicon")
output_folder = os.path.join(cwd, "profile-icons")
quality = 85

print(f"Current working directory: {cwd}")
print(f"Contents of the current directory: {os.listdir(cwd)}")

if not os.path.exists(input_folder):
    print(f"Input folder {input_folder} not found.")
else:
    batch_convert_png_to_jpg(input_folder, output_folder, quality)
    print(f"Images converted successfully and saved in {output_folder} with quality set to {quality}.")
